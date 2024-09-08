/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["web/template/*.templ", "web/template/component/*.templ", "web/template/view/*.templ"],
    theme: {
        extend: {},
    },
    plugins: [
        require('daisyui'),
    ],
    daisyui: {
        themes: false, // false: only light + dark | true: all themes | array: specific themes like this ["light", "dark", "cupcake"]
        darkTheme: "light", // name of one of the included themes for dark mode
        base: true, // applies background color and foreground color for root element by default
        styled: true, // include daisyUI colors and design decisions for all component
        utils: true, // adds responsive and modifier utility classes
        prefix: "", // prefix for daisyUI classnames (component, modifiers and responsive class names. Not colors)
        logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
        themeRoot: ":root", // The element that receives theme color CSS variables
    },
}
