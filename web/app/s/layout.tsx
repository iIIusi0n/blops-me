import {Sidebar} from "@/components/sidebar";
import {redirect} from "next/navigation";
import {getAuthorized} from "@/components/api/auth";

// @ts-ignore
export default async function Layout({ children }) {
    const authorized = await getAuthorized();

    if (!authorized) {
        redirect("/")
    }

    return (
        <div className="flex min-h-screen w-full">
            <Sidebar/>
            <div className="flex flex-1 flex-col px-8 py-10">
                {children}
            </div>
        </div>
    )
}