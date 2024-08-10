import { useEffect, useState } from "react";

import { threadGet } from "src/api/openapi-client/threads";
import { Thread } from "src/api/openapi-schema";
import { handleError } from "src/components/site/ErrorBanner";

export type Props = { editing?: string };

export function useComposeScreen({ editing }: Props) {
  const [loadingDraft, setLoadingDraft] = useState(editing !== undefined);
  const [draft, setDraft] = useState<Thread | undefined>(undefined);

  useEffect(() => {
    const getDraft = async () => {
      if (editing === undefined) return;

      const thread = await threadGet(editing);

      setDraft(thread);

      return;
    };

    getDraft()
      .catch(handleError)
      .finally(() => setLoadingDraft(false));
  }, [editing]);

  return {
    loadingDraft,
    draft,
  };
}
