import { Table, TableHeader, TableRow, TableHead, TableBody, TableCell } from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"

export function FileUploadProgress() {
  return (
    <>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">File Uploads</h1>
      </div>
      <div className="border rounded-lg overflow-hidden">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>File Name</TableHead>
              <TableHead>Type</TableHead>
              <TableHead>Progress</TableHead>
              <TableHead>Path</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell className="font-medium">image-1.jpg</TableCell>
              <TableCell>
                <Badge variant="outline" className="px-2 py-1">
                  JPEG
                </Badge>
              </TableCell>
              <TableCell>
                <div className="relative h-3 w-full overflow-hidden rounded-full bg-muted">
                  <div className="absolute h-full w-[75%] bg-primary" aria-label="75% complete" />
                </div>
              </TableCell>
              <TableCell>/uploads/image-1.jpg</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">document.pdf</TableCell>
              <TableCell>
                <Badge variant="outline" className="px-2 py-1">
                  PDF
                </Badge>
              </TableCell>
              <TableCell>
                <div className="relative h-3 w-full overflow-hidden rounded-full bg-muted">
                  <div className="absolute h-full w-full bg-primary" aria-label="100% complete" />
                </div>
              </TableCell>
              <TableCell>/uploads/document.pdf</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">video.mp4</TableCell>
              <TableCell>
                <Badge variant="outline" className="px-2 py-1">
                  MP4
                </Badge>
              </TableCell>
              <TableCell>
                <div className="relative h-3 w-full overflow-hidden rounded-full bg-muted">
                  <div className="absolute h-full w-[50%] bg-primary" aria-label="50% complete" />
                </div>
              </TableCell>
              <TableCell>/uploads/video.mp4</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">presentation.pptx</TableCell>
              <TableCell>
                <Badge variant="outline" className="px-2 py-1">
                  PPTX
                </Badge>
              </TableCell>
              <TableCell>
                <div className="relative h-3 w-full overflow-hidden rounded-full bg-muted">
                  <div className="absolute h-full w-[90%] bg-primary" aria-label="90% complete" />
                </div>
              </TableCell>
              <TableCell>/uploads/presentation.pptx</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">spreadsheet.xlsx</TableCell>
              <TableCell>
                <Badge variant="outline" className="px-2 py-1">
                  XLSX
                </Badge>
              </TableCell>
              <TableCell>
                <div className="relative h-3 w-full overflow-hidden rounded-full bg-muted">
                  <div className="absolute h-full w-[30%] bg-primary" aria-label="30% complete" />
                </div>
              </TableCell>
              <TableCell>/uploads/spreadsheet.xlsx</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>
    </>
  )
}
