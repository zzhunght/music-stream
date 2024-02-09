import React from 'react'
import { DataTable } from './data-table'
import { columns } from "./columns"
import TitlePage from '@/components/titlepage/TitlePage'
import { Pagination } from '@/components/ui/pagination'
import { Button } from '@/components/ui/button'

const data = [ 
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
  {
    no: "m5gr84i9",
    name: 316,
    artists: "success",
    duration: "ken99@yahoo.com",
    releasedate: "ken99@yahoo.com",
  },
]


function About() {
  return (
    <div>
      <TitlePage title={'Songs'}/>
      <div className="container mx-auto py-10">
        <div>
          <Button variant='outline' className='mb-10 border-gray-600'>
            Add Songs
          </Button>
        </div>
        <DataTable columns={columns} data={data} />
      </div>
    </div>
  )
}

export default About