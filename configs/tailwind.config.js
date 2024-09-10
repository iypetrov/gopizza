/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      "./templates/**/*.templ",
      "./web/js/**/*.js"
  ],
  theme: {
    extend: {
      colors: {
        border: "#ececec",
        text: "#303030",
        info: "#006aff",
        warning: "#ff0000",
        close: "#c8c8c8",
      },
    },
  },
  plugins: [],
};
