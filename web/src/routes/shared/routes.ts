import { RouteObject } from "react-router-dom";
import { RouteNames } from "../../enums";
import { Home, About } from "./pages";
import { RouteNameToPathMap } from "@constants";

export const SharedRoutes: RouteObject[] = [
  {
    id: RouteNames.Home.toString(),
    Component: Home,
    index: true
  },
  {
    id: RouteNames.About.toString(),
    Component: About,
    path: RouteNameToPathMap[RouteNames.About]
  }
];
