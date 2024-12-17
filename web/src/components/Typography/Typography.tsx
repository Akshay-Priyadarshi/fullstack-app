import { cn, StringUtils } from "@lib";
import { useCallback, useState } from "react";

type ITypographyVariants = "h1" | "h2" | "h3" | "h4" | "h5" | "h6" | "body";

type ITypographyProps = {
  variant?: ITypographyVariants;
  titleCase?: boolean;
  children: string;
  className?: string;
};

export const Typography = (props: ITypographyProps) => {
  const [variant] = useState<ITypographyVariants>(props.variant ?? "body");
  const [titleCase] = useState<boolean>(props.titleCase ?? true);

  const getFormattedText = useCallback(() => {
    return titleCase ? StringUtils.titleCase(props.children) : props.children;
  }, [props.children, titleCase]);

  const getTextVariant = useCallback(() => {
    switch (variant) {
      case "h1":
        return (
          <h1 className={cn("text-5xl", props.className)}>
            {getFormattedText()}
          </h1>
        );
      case "h2":
        return (
          <h2 className={cn("text-4xl", props.className)}>
            {getFormattedText()}
          </h2>
        );
      case "h3":
        return (
          <h3 className={cn("text-3xl", props.className)}>
            {getFormattedText()}
          </h3>
        );
      case "h4":
        return (
          <h4 className={cn("text-2xl", props.className)}>
            {getFormattedText()}
          </h4>
        );
      case "h5":
        return (
          <h5 className={cn("text-xl", props.className)}>
            {getFormattedText()}
          </h5>
        );
      case "h6":
        return (
          <h6 className={cn("text-lg", props.className)}>
            {getFormattedText()}
          </h6>
        );
      case "body":
        return (
          <p className={cn("text-base", props.className)}>
            {getFormattedText()}
          </p>
        );
      default:
        return (
          <p className={cn("text-base", props.className)}>
            {getFormattedText()}
          </p>
        );
    }
  }, [getFormattedText, variant, props.className]);

  return getTextVariant();
};
