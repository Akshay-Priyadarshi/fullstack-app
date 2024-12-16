/// <reference types="vitest/globals" />
import "@testing-library/jest-dom";

import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import en from "./src/assets/translation/en.json";

vi.mock("zustand");

beforeAll(() => {
  global.ResizeObserver = vi.fn().mockImplementation(() => ({
    observe: vi.fn(),
    unobserve: vi.fn(),
    disconnect: vi.fn()
  }));

  i18n.use(initReactI18next).init({
    lng: "en",
    fallbackLng: "en",
    resources: {
      en
    }
  });
});
