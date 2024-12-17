import { RouteObject } from "react-router-dom";
import { RouteNames } from "@enums";
import { RouteNameToPathMap } from "@constants";
import { Dashboard } from "./pages";

export const UserRoutes: RouteObject[] = [
  {
    id: RouteNames.UserDashboard.toString(),
    path: RouteNameToPathMap[RouteNames.UserDashboard],
    Component: Dashboard
  }
];
