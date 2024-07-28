import {MainPage} from "@/components/main-page";
import {redirect} from "next/navigation";
import {getAuthorized} from "@/components/api/auth";

export default async function Page() {
    const authorized = await getAuthorized();

    if (authorized) {
        redirect("/s")
    }

    return <MainPage/>;
}
