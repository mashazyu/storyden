"use client";

import { ChatBubbleLeftIcon as ChatBubbleLeftIconOutline } from "@heroicons/react/24/outline";
import { ChatBubbleLeftIcon as ChatBubbleLeftIconSolid } from "@heroicons/react/24/solid";
import { Controller, ControllerProps } from "react-hook-form";

import { Unready } from "src/components/site/Unready";

import { Byline } from "@/components/content/Byline";
import { ContentComposer } from "@/components/content/ContentComposer/ContentComposer";
import { CancelAction } from "@/components/site/Action/Cancel";
import { SaveAction } from "@/components/site/Action/Save";
import { Breadcrumbs } from "@/components/thread/Breadcrumbs";
import { ReplyBox } from "@/components/thread/ReplyBox/ReplyBox";
import { ReplyList } from "@/components/thread/ReplyList/ReplyList";
import { ThreadMenu } from "@/components/thread/ThreadMenu/ThreadMenu";
import { FormErrorText } from "@/components/ui/FormErrorText";
import { Heading } from "@/components/ui/heading";
import { HeadingInput } from "@/components/ui/heading-input";
import { HStack, LStack, styled } from "@/styled-system/jsx";

import { Form, Props, useThreadScreen } from "./useThreadScreen";

export function ThreadScreen(props: Props) {
  const { ready, error, form, isEditing, isEmpty, resetKey, data, handlers } =
    useThreadScreen(props);

  if (!ready) {
    return <Unready error={error} />;
  }

  const { thread } = data;

  return (
    <LStack gap="4">
      <styled.form
        display="flex"
        flexDirection="column"
        alignItems="start"
        gap="1"
        width="full"
        onSubmit={handlers.handleSave}
      >
        <HStack w="full" justify="space-between" alignItems="start">
          <Breadcrumbs thread={thread} />

          <HStack>
            {isEditing && (
              <>
                <CancelAction
                  type="button"
                  onClick={handlers.handleDiscardChanges}
                >
                  Discard
                </CancelAction>
                <SaveAction type="submit" disabled={isEmpty}>
                  Save
                </SaveAction>
              </>
            )}

            <ThreadMenu thread={thread} />
          </HStack>
        </HStack>

        <Byline
          href={`#${thread.id}`}
          author={thread.author}
          time={new Date(thread.createdAt)}
          updated={new Date(thread.updatedAt)}
        />

        <FormErrorText>{form.formState.errors.root?.message}</FormErrorText>

        {isEditing ? (
          <TitleInput name="title" control={form.control} />
        ) : (
          <Heading fontSize="heading.variable.1" fontWeight="bold">
            {thread.title}
          </Heading>
        )}

        <ThreadBodyInput
          control={form.control}
          name="body"
          initialValue={thread.body}
          resetKey={resetKey}
          disabled={!isEditing}
          handleEmptyStateChange={handlers.handleEmptyStateChange}
        />
      </styled.form>

      <styled.p display="flex" gap="1" alignItems="center" color="fg.muted">
        <span>
          {thread.reply_status.replied ? (
            <ChatBubbleLeftIconSolid
              width="1rem"
              title="You have replied to this thread"
            />
          ) : (
            <ChatBubbleLeftIconOutline
              width="1rem"
              title="You have not replied to this thread"
            />
          )}
        </span>
        <span>{thread.reply_status.replies} replies</span>
      </styled.p>

      <ReplyList thread={thread} />

      <ReplyBox {...thread} />
    </LStack>
  );
}

type TitleInputProps = Omit<ControllerProps<Form>, "render">;

export function TitleInput({ control }: TitleInputProps) {
  return (
    <Controller<Form>
      render={({ field: { onChange, ...field }, formState, fieldState }) => {
        return (
          <>
            <HeadingInput
              id="title-input"
              placeholder="Thread title..."
              onValueChange={onChange}
              defaultValue={formState.defaultValues?.["title"]}
              {...field}
            />

            <FormErrorText>{fieldState.error?.message}</FormErrorText>
          </>
        );
      }}
      control={control}
      name="title"
    />
  );
}

type ThreadBodyInputProps = Omit<ControllerProps<Form>, "render"> & {
  initialValue: string;
  resetKey: string;
  handleEmptyStateChange: (isEmpty: boolean) => void;
};

function ThreadBodyInput({
  control,
  name,
  initialValue,
  resetKey,
  disabled,
  handleEmptyStateChange,
}: ThreadBodyInputProps) {
  return (
    <Controller<Form>
      render={({ field: { onChange } }) => {
        function handleChange(value: string, isEmpty: boolean) {
          handleEmptyStateChange(isEmpty);
          onChange(value);
        }

        return (
          <ContentComposer
            initialValue={initialValue}
            onChange={handleChange}
            resetKey={resetKey}
            disabled={disabled}
          />
        );
      }}
      control={control}
      name={name}
    />
  );
}
