import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { RootRoute } from "@routes";
import axios from "axios";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const RootRouter = createBrowserRouter([RootRoute]);
const queryClient = new QueryClient();
axios.defaults.baseURL = import.meta.env.VITE_APP_API_BASE_URL;

function App() {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={RootRouter} />
      </QueryClientProvider>
    </>
  );
}

export default App;
