'use client';

import {UploadIcon} from "@/components/icons";
import {useCallback, useState} from 'react';
import {decodeString} from "@/components/utils/encoding";
import {useRouter} from "next/navigation";

export default function FileUpload({storageName}: { storageName: string }) {
    const [isDragging, setIsDragging] = useState(false);
    const router = useRouter();

    const handleDragEnter = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
    }) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(true);
    }, []);

    const handleDragLeave = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
    }) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(false);
    }, []);

    const handleDragOver = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
    }) => {
        e.preventDefault();
        e.stopPropagation();
    }, []);

    const findStorageIdByName = async (storageName: string) => {
        const res = await fetch(`${process.env.NEXT_PUBLIC_APP_URL}/api/storage`);
        if (!res.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await res.json();
        const storages = data.storages;
        const storage = storages.find((storage: any) => storage.name === decodeString(storageName));
        return storage.id;
    }

    const uploadFiles = async (files: any) => {
        const formData = new FormData();
        files.forEach((file: any) => {
            formData.append('files', file);
        });

        try {
            const storageId = await findStorageIdByName(storageName);

            const res = await fetch(`${process.env.NEXT_PUBLIC_APP_URL}/api/storage/${storageId}/file`, {
                method: 'POST',
                body: formData,
            });

            if (!res.ok) {
                throw new Error('Network response was not ok');
            }

            alert('Files uploading started. It may take some time to complete.');
        } catch (error) {
            console.error('An error occurred:', error);
            alert('An error occurred while uploading the files');
        }
    };

    const handleDrop = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
        dataTransfer: { files: any; };
    }) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(false);

        const files = [...e.dataTransfer.files];

        uploadFiles(files).finally(() => {
            router.push(`/s/${storageName}`);
        });
    }, [router]);

    const handleClick = useCallback(() => {
        const input = document.createElement('input');
        input.type = 'file';
        input.multiple = true;
        input.onchange = (e) => {
            // @ts-ignore
            const files = [...e.target.files];

            uploadFiles(files).finally(() => {
                router.push(`/s/${storageName}`);
            });
        };
        input.click();
    }, [router]);

    return (
        <>
            <div className="flex items-center justify-between mb-6">
                <h1 className="text-2xl font-bold">Upload files to {decodeString(storageName)}</h1>
            </div>
            <div
                className={`border rounded-lg p-6 flex items-center justify-center h-[100%] bg-muted ${isDragging ? 'border-primary' : ''}`}
                onDragEnter={handleDragEnter}
                onDragOver={handleDragOver}
                onDragLeave={handleDragLeave}
                onDrop={handleDrop}
                onClick={handleClick}
            >
                <div className="text-center space-y-2">
                    <UploadIcon className="w-12 h-12 text-primary mx-auto"/>
                    <p className="text-lg font-medium">Drag and drop files to
                        upload</p>
                    <p className="text-muted-foreground">or click to select
                        files from your device</p>
                </div>
            </div>
        </>
    )
}
