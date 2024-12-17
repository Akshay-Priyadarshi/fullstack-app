import { RouteObject } from "react-router-dom";
import { RouteNames } from "@enums";
import { Root } from "./Root";
import { UserRoutes } from "../personas";
import { SharedRoutes } from "../shared";

export const RootRoute: RouteObject = {
  id: RouteNames.Root.toString(),
  path: "/",
  Component: Root,
  children: [...SharedRoutes, ...UserRoutes]
};
