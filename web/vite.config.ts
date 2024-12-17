import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
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
  }
});
