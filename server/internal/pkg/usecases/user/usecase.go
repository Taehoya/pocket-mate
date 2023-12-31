package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	pathutil "github.com/Taehoya/pocket-mate/internal/pkg/utils/path"
	"github.com/Taehoya/pocket-mate/internal/pkg/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type NicknameSource struct {
	Adj  []string `json:"adj"`
	Name struct {
		Animal []string `json:"animal"`
	} `json:"name"`
}

type UserUseCase struct {
	userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) Get(ctx context.Context, id int) (*entities.User, error) {
	user, err := u.userRepository.GetUserById(ctx, id)

	return user, err
}

func (u *UserUseCase) Register(ctx context.Context, email, password string) (*entities.User, error) {
	user, err := u.userRepository.GetUser(ctx, email)

	if err != nil {
		return nil, err
	}

	if user.Email() != "" {
		return nil, fmt.Errorf("duplicated identification")
	}

	hashedPasword, err := encrpyt(password)
	if err != nil {
		return nil, err
	}

	nickname, err := getNickNameFromSource()
	if err != nil {
		return nil, err
	}

	user, err = u.userRepository.SaveUser(ctx, email, hashedPasword, nickname)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepository.GetUser(ctx, email)

	if user == nil || err != nil {
		return "", fmt.Errorf("invalid identification")
	}

	err = user.CheckPassword(password)
	if err != nil {
		return "", fmt.Errorf("invalid identification")
	}

	token, err := token.MakeToken(user.ID())
	if err != nil {
		log.Printf("unable to make token: %v\n", err)
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

func encrpyt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to encrypt password: %v\n", err)
		return "", fmt.Errorf("internal server error")
	}

	return string(hash), nil
}

func getNickNameFromSource() (string, error) {
	rand.Seed(time.Now().UnixNano())

	rootPath, err := pathutil.GetProjectRootDir()
	if err != nil {
		log.Printf("failed to get project root directory: %v\n", err)
		return "", fmt.Errorf("internal server error")
	}

	sourceFile := path.Join(rootPath, "internal/pkg", "resources", "nickname.json")
	file, err := os.Open(sourceFile)
	if err != nil {
		log.Printf("failed to open file: %v\n", err)
		return "", fmt.Errorf("internal server error")
	}
	defer file.Close()

	var source NicknameSource
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&source); err != nil {
		log.Printf("failed to decode json: %v\n", err)
		return "", fmt.Errorf("internal server error")
	}

	randomAdj := source.Adj[rand.Intn(len(source.Adj))]
	randomAnimal := source.Name.Animal[rand.Intn(len(source.Name.Animal))]

	nickname := fmt.Sprintf("%s %s", randomAdj, randomAnimal)

	return nickname, nil
}
