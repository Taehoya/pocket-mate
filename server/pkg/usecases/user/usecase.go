package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	pathutil "github.com/Taehoya/pocket-mate/pkg/utils/path"
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
	hashedPasword, err := encrpyt(password)
	if err != nil {
		return nil, err
	}

	nickname, err := getNickNameFromSource()
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.SaveUser(ctx, email, hashedPasword, nickname)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) Login(ctx context.Context, email, password string) {

}

func encrpyt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generated password")
	}

	return string(hash), nil
}

func getNickNameFromSource() (string, error) {
	rand.Seed(time.Now().UnixNano())

	rootPath, err := pathutil.GetRootPath()
	if err != nil {
		return "", fmt.Errorf("failed to get root path: %v", err)
	}

	sourceFile := path.Join(rootPath, "pkg", "resources", "nickname.json")
	file, err := os.Open(sourceFile)
	if err != nil {
		return "", fmt.Errorf("failed to open source: %v", err)
	}
	defer file.Close()

	var source NicknameSource
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&source); err != nil {
		return "", fmt.Errorf("failed to decode json: %v", err)
	}

	randomAdj := source.Adj[rand.Intn(len(source.Adj))]
	randomAnimal := source.Name.Animal[rand.Intn(len(source.Name.Animal))]

	nickname := fmt.Sprintf("%s %s", randomAdj, randomAnimal)

	return nickname, nil
}
