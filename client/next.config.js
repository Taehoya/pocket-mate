/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "standalone",
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: `${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/:path*`,
      },
    ];
  },
  images: {
    domains: ["flagsapi.com"],
  },
};

module.exports = nextConfig;
