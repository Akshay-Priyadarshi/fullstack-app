import { RouteNames } from "../enums/RouteNames";

export const RouteNameToPathMap: Record<RouteNames, string> = {
  [RouteNames.Root]: "/",
  [RouteNames.Home]: "/",
  [RouteNames.About]: "/about",
  [RouteNames.UserDashboard]: "/user/dashboard"
};
