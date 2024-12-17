import { Navbar } from "@components";
import { useTranslation } from "react-i18next";
import { Typography } from "@components";

export const Home = () => {
  const { t } = useTranslation();
  return (
    <main className="p-16">
      <Navbar />
      <Typography>{t("greeting")}</Typography>
    </main>
  );
};
