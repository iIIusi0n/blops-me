import {FileExplorer} from "@/components/file-explorer";
import {EmptyStorage} from "@/components/empty-storage";

function getData() {
  const files = [
    {name: 'file1', type: 'JPEG', modifiedAt: '2021-01-01', size: "21 MB"},
    {name: 'file2', type: 'PPTX', modifiedAt: '2021-01-01', size: "21 MB"},
    {name: 'file3', type: 'MP3', modifiedAt: '2021-01-01', size: "21 MB"},
  ]

  return files;
}

export default async function Page() {
  return <EmptyStorage />;
}
