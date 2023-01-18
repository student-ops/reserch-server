import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

async function insertMedicenData() {
    const Users = await prisma.user.findMany()
    console.log(Users)
    await prisma.med_schedule.create({
        data: {
            user_id: 1,
            state: true,
            taken_day: 30,
            taken_time: 1,
            name: "成長ホルモン",
        },
    })
    const allUsers = await prisma.med_schedule.findMany({})
    console.dir(allUsers, { depth: null })
    await prisma.$disconnect()
}
insertMedicenData().catch(async (e) => {
    console.error(e)
    await prisma.$disconnect()
    process.exit(1)
})

async function insertUser() {
    await prisma.user.create({
        data: {
            name: "john",
        },
    })
    await prisma.$disconnect()
}
insertUser().catch(async (e) => {
    console.error(e)
    await prisma.$disconnect()
    process.exit(1)
})
