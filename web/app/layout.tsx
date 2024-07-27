import { IBM_Plex_Mono } from 'next/font/google'
import { cn } from '@/lib/utils'
import './globals.css'

const fontHeading = IBM_Plex_Mono({
  weight: "700",
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-heading'
})

const fontBody = IBM_Plex_Mono({
  weight: "400",
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-body',
})

// @ts-ignore
export default function Layout({ children }) {
  return (
      <html lang="en">
      <head>
          <link rel="icon" href="/app/favicon.ico" sizes="any"/>
          <title>Blops Me</title>
      </head>
      <body
          className={cn(
              'antialiased',
              fontHeading.variable,
              fontBody.variable
          )}
      >
      {children}
      </body>
      </html>
  )
}