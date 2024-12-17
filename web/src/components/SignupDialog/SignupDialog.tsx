import { useForm, SubmitHandler } from "react-hook-form";
import {
  Button,
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  Input
} from "@ui";
import { Typography } from "../Typography";
import { useTranslation } from "react-i18next";

type SignupFormInputs = {
  email: string;
  password: string;
};

export const SignupDialog = () => {
  const { t } = useTranslation();
  const { register, handleSubmit } = useForm<SignupFormInputs>();
  const onLoginFormSubmit: SubmitHandler<SignupFormInputs> = () => {};
  return (
    <Dialog>
      <DialogTrigger>
        <Button>
          <Typography>{t("menuitems.signup")}</Typography>
        </Button>
      </DialogTrigger>
      <DialogContent className="p-8">
        <DialogHeader>
          <DialogTitle className="m-auto">
            Create your Auth24 account
          </DialogTitle>
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
