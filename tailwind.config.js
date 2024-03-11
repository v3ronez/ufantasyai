/** @type {import('tailwindcss').Config} */
module.exports = {
  // content: ["./view/**/*.html", "./**/*.templ", "./**/*.go",],
  content: ["./view/**/*.templ", "./**/*.templ"],
  safelist: [],
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["dark"]
  }
}