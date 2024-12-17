import { Link } from "react-router-dom";
import { LoginDialog } from "../LoginDialog";
import { SignupDialog } from "../SignupDialog";
import { LanguageSelector } from "../LanguageSelector/LanguageSelector";
import { Typography } from "../Typography";

export interface IMenubarProps {
  menuLinks: { label: string; path: string }[];
}

export const Menubar: React.FC<IMenubarProps> = ({ menuLinks }) => {
  return (
    <menu className="flex gap-8 justify-between items-center">
      <div className="flex gap-6">
        {menuLinks.map((menuLink: { label: string; path: string }) => {
          return (
            <li key={menuLink.path}>
              <Link to={menuLink.path}>
                <Typography>{menuLink.label}</Typography>
              </Link>
            </li>
          );
        })}
      </div>
      <div className="flex gap-6">
        <LoginDialog />
        <SignupDialog />
      </div>
      <LanguageSelector />
    </menu>
  );
};
