import { PrismaClient } from "@prisma/client"
const prisma = new PrismaClient()

let everyday: number = (1 << 7) - 1
let onry_wed: number = 1 << 4
let sun_mon_wed: number = 1 + 2 + (1 << 4)
const med_data = [
    {
        user_id: 1,
        state: true,
        taken_day: everyday,
        taken_time: 1,
        name: "シプロキサン",
        tabletes: 2,
    },
    {
        user_id: 1,
        state: true,
        taken_day: sun_mon_wed,
        taken_time: 2,
        name: "クラビット",
        tabletes: 1,
    },
    {
        user_id: 1,
        state: true,
        taken_day: onry_wed,
        taken_time: 2,
        name: "ステロイド",
        tabletes: 3,
    },
    {
        user_id: 1,
        state: true,
        taken_day: everyday,
        taken_time: 0,
        name: "メジコン",
        tabletes: 2,
    },
    {
        user_id: 1,
        state: true,
        taken_day: everyday,
        taken_time: 0,
        name: "ロキソマリン",
        tabletes: 2,
    },
]
async function insertMedicenData() {
    await prisma.med_schedule.createMany({
        data: med_data,
    })
}
async function insertUser() {
    await prisma.user.createMany({
        data: [
            {
                name: "john",
            },
            {
                name: "sakamoto",
            },
            {
                name: "messi",
            },
        ],
    })
}

insertMedicenData()
    .then(async () => {
        console.log("inserted med-data")
        await prisma.$disconnect()
    })
    .catch(async (e) => {
        console.error(e)
        await prisma.$disconnect()
        process.exit(1)
    })
insertUser()
    .then(async () => {
        await prisma.$disconnect()
    })
    .catch(async (e) => {
        console.error(e)
        await prisma.$disconnect()
        process.exit(1)
    })
