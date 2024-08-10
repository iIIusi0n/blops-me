import {cookies} from "next/headers";
import {getStorages} from "@/components/api/storage";

export interface File {
    id: number;
    name: string;
    type: string;
    last_modified: string;
    size: number;
}

export async function resolveStorageID(storageName: string): Promise<number> {
    const data = await getStorages();
    const storage = data.storages.find(storage => storage.name === storageName);
    return storage?.id ?? -1;
}

export async function getFiles(storageId: number): Promise<File[]> {
    const token = cookies().get('token')?.value;
    const resp = await fetch(`${process.env.APP_API_URL}/api/storage/${storageId}/file`, {
        credentials: 'include',
        headers: {
            'Cookie': `token=${token}`
        }
    });
    if (!resp.ok) {
        return [];
    }

    return resp.json();
}
