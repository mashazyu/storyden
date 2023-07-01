import { useFormContext } from "react-hook-form";

import { ThreadCreate } from "../ComposeForm/useComposeForm";

export function useTitleInput() {
  const ctx = useFormContext<ThreadCreate>();

  return {
    control: ctx.control,
    fieldError: ctx.formState.errors.title,
  };
}