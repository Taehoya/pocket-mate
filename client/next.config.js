/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: "http://54.85.236.192/api/:path*",
      },
    ];
  },
};

module.exports = nextConfig;
