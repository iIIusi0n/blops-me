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
            {children}
        </div>
    )
}
