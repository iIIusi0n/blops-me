import FileUpload from "@/components/file-upload";

export default function Page({params}: { params: { storage: string } }) {
    return <FileUpload storageName={params.storage}/>;
}
