import { StoryContext, StoryFn } from "@storybook/react";
import { createMemoryRouter, RouterProvider } from "react-router-dom";
import { RouteNames } from "@enums";
import { RouteNameToPathMap } from "@constants";
import { Error } from "@components";

export const withRouter = (Story: StoryFn, StoryCtx: StoryContext) => {
  const { routeId, routePath, data } = StoryCtx;
  const initialRoutePath = routePath ?? RouteNameToPathMap[RouteNames.Root];
  const memoryRouter = createMemoryRouter(
    [
      {
        id: routeId ?? RouteNames.Root,
        index: initialRoutePath,
        element: <Story />,
        errorElement: <Error text="" />,
        loader: () => {
          return { ...data };
        }
      }
    ],
    { initialEntries: [initialRoutePath] }
  );
  return <RouterProvider router={memoryRouter} />;
};
