import {Sidebar} from "@/components/sidebar";
import {EmptyStorage} from "@/components/empty-storage";

export default function Page() {
    return (
        <>
            <Sidebar encodedStorageName="" />
            <div className="flex flex-1 flex-col px-8 py-10">
                <EmptyStorage/>
            </div>
        </>
    );
}
