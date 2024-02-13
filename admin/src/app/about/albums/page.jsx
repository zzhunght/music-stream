import TitlePage from '@/components/titlepage/TitlePage'
import { Button } from '@/components/ui/button'
import React from 'react'
import { DataTable } from './data-table'
import { columns } from './columns'

const data = [ 
  {
    no: "m5gr84i9",
    name: 316,
    songs: "ken99@yahoo.com",
    artists: "success",
    createdat: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    songs: "ken99@yahoo.com",
    artists: "success",
    createdat: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    songs: "ken99@yahoo.com",
    artists: "success",
    createdat: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    songs: "ken99@yahoo.com",
    artists: "success",
    createdat: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    songs: "ken99@yahoo.com",
    artists: "success",
    createdat: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    songs: "ken99@yahoo.com",
    artists: "success",
    createdat: "ken99@yahoo.com",
  },
]

function Albums() {
  return (
    <div>
      <TitlePage title={'Albums'}/>
      <div className="container mx-auto py-10">
        <div>
          <Button variant='outline' className='mb-10 border-gray-600'>
            Add Albums
          </Button>
        </div>
        <DataTable columns={columns} data={data} />
      </div>
    </div>
  )
}

export default Albums;