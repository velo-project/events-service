using System;
using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Velo.EventsService.Persistence.Migrations
{
    /// <inheritdoc />
    public partial class CreateBaseEventsTable : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "tb_events",
                columns: table => new
                {
                    id_event = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    name_event = table.Column<string>(type: "text", nullable: false),
                    description_event = table.Column<string>(type: "text", nullable: true),
                    photo_event = table.Column<string>(type: "text", nullable: true),
                    active_event = table.Column<bool>(type: "boolean", nullable: false),
                    canceled_event = table.Column<bool>(type: "boolean", nullable: false),
                    when_will_happen_event = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    created_at = table.Column<DateTime>(type: "timestamp with time zone", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_tb_events", x => x.id_event);
                });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "tb_events");
        }
    }
}
