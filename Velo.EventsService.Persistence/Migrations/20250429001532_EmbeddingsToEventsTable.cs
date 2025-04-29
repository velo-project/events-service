using Microsoft.EntityFrameworkCore.Migrations;
using Pgvector;

#nullable disable

namespace Velo.EventsService.Persistence.Migrations
{
    /// <inheritdoc />
    public partial class EmbeddingsToEventsTable : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AlterDatabase()
                .Annotation("Npgsql:PostgresExtension:vector", ",,");

            migrationBuilder.AddColumn<Vector>(
                name: "embeddings_event",
                table: "tb_events",
                type: "vector(3072)",
                nullable: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "embeddings_event",
                table: "tb_events");

            migrationBuilder.AlterDatabase()
                .OldAnnotation("Npgsql:PostgresExtension:vector", ",,");
        }
    }
}
