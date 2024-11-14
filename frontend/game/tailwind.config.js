export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
colors: {
        "dark-blue": "#212a31",
        "medium-blue": "#2e3944",
        "bright-blue": "#124e66",
        "soft-blue": "#748d92",
        "light-gray": "#d3d9d4",
      },
      animation: {
        fadeIn: 'fadeIn 0.5s ease-in-out',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
      },
    },
  },
  plugins: [],
}

