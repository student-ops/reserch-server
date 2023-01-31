import { med_schedule, PrismaClient } from "@prisma/client"
import { time } from "console"
const prisma = new PrismaClient()

export async function selectAll() {
    const fetched = await prisma.med_schedule.findMany({})
    console.log(fetched)
    return fetched
}

async function selectMed(id: number) {
    const fetched = await prisma.med_schedule.findMany({
        where: {
            state: true,
            user_id: id,
        },
    })
    prisma.$disconnect()
    return fetched
}
function medcineQuote(
    fetched: med_schedule[],
    date: number,
    time: number
): string {
    let quote: string
    switch (time) {
        case 0:
            quote = "朝のお薬は"
            break
        case 1:
            quote = "お昼のお薬は"
            break
        case 2:
            quote = "夜のお薬は"
            break
        default:
            quote = ""
            break
    }
    let sentence: string = ""
    //flug　リファクタリングできそう
    let flug: number = 0
    for (let i = 0; i < fetched.length; i++) {
        if ((fetched[i].taken_day & date) / date == 1) {
            if (fetched[i].taken_time == time) {
                sentence = fetched[i].name + "と"
                quote += sentence
                flug++
            }
        }
    }
    if (flug > 0) {
        quote = quote.slice(0, -1)
        quote = quote.concat("です")
    } else {
        quote += "ありません"
    }
    return quote
}
export async function getMessage(id: number) {
    let now = new Date()
    let now_h = now.getHours() + (now.getTimezoneOffset() + 540) / -60
    let date = 1 << new Date().getDay()
    let time_zone: number
    if (now_h < 8) {
        time_zone = 0
    } else if (now_h < 13) {
        time_zone = 1
    } else if (now_h < 21) {
        time_zone = 2
    } else {
        date << 1
        time_zone = 0
    }
    let fetched = await selectMed(id)
    console.log("date = " + date + "time = " + now_h)
    let message = medcineQuote(fetched, date, time_zone)
    console.log(message)
    return message
}
