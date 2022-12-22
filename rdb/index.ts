import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

async function main() {
    await prisma.initial.create({
        data: {
            name: "坂本琉太",
            age: 22,
        },
    })

    const allUsers = await prisma.initial.findMany({})
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
