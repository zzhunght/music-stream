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
    accessorKey: "songs",
    header: "Songs",
  },
  {
    accessorKey: "artists",
    header: "Artists",
  },
  {
    accessorKey: "createdat",
    header: "Created At",
  },
]
