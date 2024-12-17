import { Link } from "react-router-dom";
import { IMenubarProps, Menubar } from "../Menubar";
import { useTranslation } from "react-i18next";
import { useMemo } from "react";

export const Navbar = () => {
  const { t } = useTranslation();
  const menuLinks = useMemo<IMenubarProps["menuLinks"]>(() => {
    return [
      {
        label: t("menuitems.about"),
        path: "/about"
      },
      {
        label: t("menuitems.dashboard"),
        path: "/user/dashboard"
      }
    ];
  }, [t]);

  return (
    <nav className="flex justify-between items-center mb-16">
      <Link to={"/"}>
        <h1 className="font-light text-3xl tracking-wide">LOGO</h1>
      </Link>
      <Menubar menuLinks={menuLinks} />
    </nav>
  );
};
