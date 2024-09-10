/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      "./templates/**/*.templ",
      "./web/js/**/*.js"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
  ],
  daisyui: {
    themes: [
        {
          cmyk: {
            ...require("daisyui/src/theming/themes")["cmyk"],
            success: "green",
          },
        },
    ],
    logs: true,
  },
};
