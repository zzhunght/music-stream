import React from 'react'
import { DataTable } from './data-table'
import { columns } from "./columns"
import TitlePage from '@/components/titlepage/TitlePage'

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
]


function About() {
  return (
    <div>
      <TitlePage title={'Songs'}/>
      <div className="container mx-auto py-10">
        <DataTable columns={columns} data={data} />
      </div>
    </div>
  )
}

export default About