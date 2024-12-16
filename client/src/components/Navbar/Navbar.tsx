import { Link } from "react-router-dom";
import { Menubar } from "../Menubar";

export const Navbar = () => {
  const menuLinks = [
    {
      label: "About",
      path: "/about"
    },
    {
      label: "Dashboard",
      path: "/user/dashboard"
    }
  ];

  return (
    <nav className="flex justify-between items-center mb-16">
      <Link to={"/"}>
        <h1 className="font-light text-3xl tracking-wide">LOGO</h1>
      </Link>
      <Menubar menuLinks={menuLinks} />
    </nav>
  );
};
