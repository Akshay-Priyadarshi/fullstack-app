import { Navbar } from "@components";
import { useTranslation } from "react-i18next";

export const Home = () => {
  const { t } = useTranslation();
  return (
    <main className="p-16">
      <Navbar />
      <h1>Home Page</h1>
      <h2>{t("greeting")}</h2>
    </main>
  );
};
