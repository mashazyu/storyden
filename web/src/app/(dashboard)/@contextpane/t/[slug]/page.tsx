import { threadGet } from "@/api/openapi-server/threads";
import { ThreadScreenContextPane } from "@/screens/thread/ThreadScreen/ThreadScreenContextPane";

export default async function Page(props: { params: { slug: string } }) {
  const { slug } = props.params;

  try {
    const { data } = await threadGet(slug);

    return <ThreadScreenContextPane slug={slug} thread={data} />;
  } catch (e) {
    console.error(e);
    return null;
  }
}