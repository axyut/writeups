import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

// https://astro.build/config
export default defineConfig({
    site: "https://axyut.github.io",
    base: "/writeups",
    integrations: [
        starlight({
            favicon: "./src/assets/logo.jpg",
            title: "writeups",
            logo: {
                src: "./src/assets/logo.jpg",
            },
            social: {
                github: "https://github.com/axyut/writeups",
            },
            sidebar: [
                {
                    label: "[home] Home",
                    link: "/",
                },
                {
                    label: "[box] Random",
                    autogenerate: {
                        directory: "random",
                    },
                },
                {
                    label: "[box] Base",
                    autogenerate: {
                        directory: "base",
                    },
                },
                {
                    label: "[box] Pwns",
                    autogenerate: {
                        directory: "pwns",
                    },
                },
                {
                    label: "[box] CTFs",
                    autogenerate: {
                        directory: "ctfs",
                    },
                },
                {
                    label: "[box] Non-Tech",
                    autogenerate: {
                        directory: "non-tech",
                    },
                },
            ],
            components: {
                ThemeProvider: "./src/components/ThemeProvider.astro",
                ThemeSelect: "./src/components/ThemeSelect.astro",
                SiteTitle: "./src/components/SiteTitle.astro",
                Sidebar: "./src/components/Sidebar.astro",
                Pagination: "./src/components/Pagination.astro",
                Hero: "./src/components/Hero.astro",
            },
            customCss: [
                "@fontsource-variable/space-grotesk/index.css",
                "@fontsource/space-mono/400.css",
                "@fontsource/space-mono/700.css",
                "./src/styles/theme.css",
            ],
            expressiveCode: {
                themes: ["github-dark"],
            },
            pagination: false,
            lastUpdated: true,
        }),
    ],
    output: "static",
});
