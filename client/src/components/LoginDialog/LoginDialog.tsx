import { SubmitHandler, useForm } from "react-hook-form";
import {
  Button,
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  Input
} from "@ui";

type LoginFormInputs = {
  email: string;
  password: string;
};

export const LoginDialog = () => {
  const { register, handleSubmit } = useForm<LoginFormInputs>();
  const onLoginFormSubmit: SubmitHandler<LoginFormInputs> = () => {};
  return (
    <Dialog>
      <DialogTrigger>
        <Button>Login</Button>
      </DialogTrigger>
      <DialogContent className="p-8">
        <DialogHeader>
          <DialogTitle className="m-auto">Login to Auth24</DialogTitle>
        </DialogHeader>
        <form
          onSubmit={handleSubmit(onLoginFormSubmit)}
          className="flex flex-col gap-4"
        >
          <Input
            type="email"
            placeholder="Email"
            {...register("email")}
          />
          <Input
            type="password"
            placeholder="Password"
            {...register("password")}
          />
          <Button type="submit">Login</Button>
        </form>
        <p className="m-auto text-gray-600">New to Auth24, Signup Here!</p>
      </DialogContent>
    </Dialog>
  );
};
