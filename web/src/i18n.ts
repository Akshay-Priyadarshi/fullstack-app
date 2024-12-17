import i18n from "i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";
import en from "./assets/translation/en.json";
import hi from "./assets/translation/hi.json";

i18n.use(LanguageDetector).use(initReactI18next).init({
  lng: "en",
  resources: {
    en,
    hi
  }
});
