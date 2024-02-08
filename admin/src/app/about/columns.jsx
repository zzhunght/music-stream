"use client"

// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.

export const columns = [
  {
    accessorKey: "no",
    header: "No",
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "artists",
    header: "Artists",
  },
  {
    accessorKey: "duration",
    header: "Duration",
  },
  {
    accessorKey: "releasedate",
    header: "Release date",
  },
]
