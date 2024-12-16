import { AppLanguages } from "@constants";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue
} from "@ui";
import { useTranslation } from "react-i18next";

export const LanguageSelector = () => {
  const { i18n } = useTranslation();

  const changeLanguage = (langKey: string) => {
    i18n.changeLanguage(langKey);
  };

  const languageSelectOptions = Object.keys(AppLanguages).map(
    (langKey: string) => {
      return (
        <SelectItem
          key={langKey}
          value={langKey}
        >
          {AppLanguages[langKey]}
        </SelectItem>
      );
    }
  );

  return (
    <div>
      <Select
        defaultValue={i18n.language}
        onValueChange={(selectedLanguageKey: string) =>
          changeLanguage(selectedLanguageKey)
        }
      >
        <SelectTrigger
          className="w-24"
          data-testid="language-selector-trigger"
        >
          <SelectValue placeholder="Language" />
        </SelectTrigger>
        <SelectContent>{languageSelectOptions}</SelectContent>
      </Select>
    </div>
  );
};
