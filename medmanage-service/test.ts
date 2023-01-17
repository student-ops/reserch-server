import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

async function main() {
    // ... you will write your Prisma Client queries here
    const Users = await prisma.user.findMany()
    console.log(Users)
    await prisma.med_schedule.create({
        data: {
            user_id: 1,
            state: true,
            name: "ステロイド",
        },
    })
    const allUsers = await prisma.med_schedule.findMany({})
    console.dir(allUsers, { depth: null })
}

main()
    .then(async () => {
        await prisma.$disconnect()
    })
    .catch(async (e) => {
        console.error(e)
        await prisma.$disconnect()
        process.exit(1)
    })

// async function main() {
//     await prisma.user.create({
//         data: {
//             name: "sakamoto",
//         },
//     })

//     const allUsers = await prisma.user.findMany({})
//     console.dir(allUsers, { depth: null })
// }
