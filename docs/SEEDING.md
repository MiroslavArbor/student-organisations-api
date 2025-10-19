# Database Seeding Guide

This guide explains how to use the SeedManager to populate your database with test data.

## 🌱 What Gets Seeded

The seeding process creates realistic test data for:

- **5 Organizations**: SSFON, Fonis, Aisec, Estiem, SportFON
- **9 Roles**: Admin, President, Vice President, Secretary, Treasurer, Project Manager, Team Lead, Member, Guest
- **16 Members**: 3-4 members per organization with hashed passwords
- **12 Teams**: 2-3 teams per organization (Development, Marketing, Events, etc.)
- **11 Projects**: Active projects with realistic timelines
- **Role Assignments**: Presidents, VPs, and members with appropriate roles
- **Team Memberships**: Members assigned to teams within their organization
- **Project Assignments**: Members assigned to projects with roles and grades

## 🚀 How to Run Seeding

### Option 1: Automatic Seeding (Recommended for Development)

Set your environment to development in `.env`:
```env
ENV=development
```

Then run your application:
```bash
go run cmd/api/main.go
```

Seeding will run automatically on startup and only insert data if tables are empty.

### Option 2: Manual Seeding Command

Run the dedicated seeding command:
```bash
go run cmd/seed/main.go
```

This will seed the database immediately without starting the main application.

### Option 3: Programmatic Usage

```go
// Initialize database connection
db := yourDatabaseConnection()

// Initialize repositories
orgRepo := repositories.NewOrganisationRepository(db)
roleRepo := repositories.NewRoleRepository(db)
memberRepo := repositories.NewMemberRepository(db)
// ... other repositories

// Initialize seeders
orgSeeder := seeding.NewOrganisationSeeder(*orgRepo)
roleSeeder := seeding.NewRoleSeeder(*roleRepo)
// ... other seeders

// Create and run seed manager
seedManager := seeding.NewSeedManager(
    orgSeeder,
    roleSeeder,
    memberSeeder,
    teamSeeder,
    projectSeeder,
    userRoleSeeder,
    teamMemberSeeder,
    memberProjectSeeder,
)

if err := seedManager.SeedAll(); err != nil {
    log.Fatal("Seeding failed:", err)
}
```

## 📋 Seeding Order

**Important**: Seeders must run in this specific order due to foreign key dependencies:

1. **OrganisationSeeder** - Creates base organizations
2. **RoleSeeder** - Creates user roles
3. **MemberSeeder** - Creates members (depends on organizations)
4. **TeamSeeder** - Creates teams (depends on organizations)
5. **ProjectSeeder** - Creates projects (depends on organizations)
6. **UserRoleSeeder** - Assigns roles to members (depends on members + roles)
7. **TeamMemberSeeder** - Assigns members to teams (depends on teams + members)
8. **MemberProjectSeeder** - Assigns members to projects (depends on projects + members)

## 🔒 Idempotent Seeding

All seeders are **idempotent** - they check if data already exists before inserting:

```go
// Example from OrganisationSeeder
existingOrgs, err := s.repo.ListAll()
if err != nil {
    return err
}
if len(existingOrgs) > 0 {
    return nil // Skip seeding if data already exists
}
```

This means you can safely run seeding multiple times without duplicating data.

## 🧪 Test Data Details

### Sample Organizations
- **SSFON**: "Best organisation in FON"
- **Fonis**: "Second best organisation in FON"
- **Aisec**: "Weird people"
- **Estiem**: "Only woman club"
- **SportFON**: "Sports organisation"

### Sample Members
All members have the password: `password123` (bcrypt hashed)

Examples:
- marko.petrovic@ssfon.rs (SSFON President)
- ana.nikolic@ssfon.rs (SSFON Vice President)
- petar.milic@fonis.rs (Fonis President)

### Sample Projects
- FON Mobile App (SSFON)
- Tech Conference 2025 (SSFON)
- Digital Library (Fonis)
- Global Exchange Program (Aisec)
- FON Olympics (SportFON)

## 🛠 Adding New Seeders

To add a new seeder:

1. Create the seeder struct implementing `Seeder` interface
2. Add appropriate checks for existing data
3. Add the seeder to `SeedManager` in correct dependency order
4. Update this documentation

Example:
```go
type NewSeeder struct {
    repo repositories.NewRepository
}

func (s *NewSeeder) Seed() error {
    // Check if data exists
    existing, err := s.repo.ListAll()
    if err != nil {
        return err
    }
    if len(existing) > 0 {
        return nil
    }
    
    // Create and insert data
    // ...
    
    return nil
}
```

## 🔧 Troubleshooting

### "Failed to seed database" error
- Check database connection
- Ensure tables are migrated (`AutoMigrate` ran)
- Verify foreign key constraints

### Duplicate data
- Seeders are idempotent, but if you modify the check logic, you might get duplicates
- Clear database and re-run migrations + seeding

### Dependency errors
- Ensure seeding order is correct
- Organizations and Roles must be seeded before Members
- Members must exist before UserRoles, TeamMembers, or MemberProjects