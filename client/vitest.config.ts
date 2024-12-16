import { defineConfig } from "vitest/config";
import path from "path";

export default defineConfig({
  resolve: {
    alias: {
      "@components": path.resolve(__dirname, "./src/components/index"),
      "@ui": path.resolve(__dirname, "./src/components/ui/index"),
      "@hooks": path.resolve(__dirname, "./src/hooks/index"),
      "@enums": path.resolve(__dirname, "./src/enums/index"),
      "@constants": path.resolve(__dirname, "./src/constants/index"),
      "@routes": path.resolve(__dirname, "./src/routes/index"),
      "@lib": path.resolve(__dirname, "./src/lib/index"),
      "@decorators": path.resolve(__dirname, "./src/decorators/index"),
      "@stores": path.resolve(__dirname, "./src/stores/index")
    }
  },
  test: {
    coverage: {
      exclude: [
        "vitest.setup.ts",
        ".storybook/**",
        "./src/mocks/*",
        "src/test-utilities.tsx",
        "**/index.ts",
        "**/*.stories.tsx",
        "**/*.test.tsx",
        "**/*.test.ts"
      ],
      reporter: ["text-summary"]
    },
    environment: "jsdom",
    globals: true,
    setupFiles: "vitest.setup.ts",
    pool: "forks",
    hookTimeout: 20000,
    testTimeout: 10000
  }
});
