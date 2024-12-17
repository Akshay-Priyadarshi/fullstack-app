import { AppLanguages } from "@enums";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue
} from "@ui";
import { useCallback, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Typography } from "../Typography";

export const LanguageSelector = () => {
  const { i18n, t } = useTranslation();

  const changeLanguage = useCallback(
    (langKey: string) => {
      i18n.changeLanguage(langKey);
    },
    [i18n]
  );

  const languageLabels = useMemo(() => {
    return {
      [AppLanguages.en]: t("languages.english"),
      [AppLanguages.hi]: t("languages.hindi")
    } as Record<AppLanguages, string>;
  }, [t]);

  const languageSelectOptions = useMemo(
    () =>
      Object.keys(AppLanguages).map((langKey: string) => {
        return (
          <SelectItem
            key={langKey}
            value={langKey}
          >
            <Typography>
              {
                languageLabels[
                  AppLanguages[langKey as keyof typeof AppLanguages]
                ]
              }
            </Typography>
          </SelectItem>
        );
      }),
    [languageLabels]
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
          className="w-26"
          data-testid="language-selector-trigger"
        >
          <SelectValue placeholder="Language" />
        </SelectTrigger>
        <SelectContent>{languageSelectOptions}</SelectContent>
      </Select>
    </div>
  );
};
