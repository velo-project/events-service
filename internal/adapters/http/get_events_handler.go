package http

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type GetEventsHandler struct {
	getRecommendedEventsService     ports.GetRecommendedEventsService
	getTrendingEventsService        ports.GetTrendingEventsService
	getLastParticipatedEventsService ports.GetLastParticipatedEventsService
	getSubscribedEventsService      ports.GetSubscribedEventsService
}

func NewGetEventsHandler(
	getRecommendedEventsService ports.GetRecommendedEventsService,
	getTrendingEventsService ports.GetTrendingEventsService,
	getLastParticipatedEventsService ports.GetLastParticipatedEventsService,
	getSubscribedEventsService ports.GetSubscribedEventsService,
) *GetEventsHandler {
	return &GetEventsHandler{
		getRecommendedEventsService:     getRecommendedEventsService,
		getTrendingEventsService:        getTrendingEventsService,
		getLastParticipatedEventsService: getLastParticipatedEventsService,
		getSubscribedEventsService:      getSubscribedEventsService,
	}
}

type GetEventsResponse struct {
	RecommendedEvents     []entities.Event `json:"recommended_events"`
	TrendingEvents        []entities.Event `json:"trending_events"`
	LastParticipatedEvents []entities.Event `json:"last_participated_events"`
	SubscribedEvents      []entities.Event `json:"subscribed_events"`
}

// @Summary Get events
// @Description Get recommended, trending, last participated and subscribed events for the authenticated user.
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} GetEventsResponse
// @Failure 500 {object} ErrorResponse
// @Router /events [get]
// @Security Bearer
func (h *GetEventsHandler) Handle(c *gin.Context) {
	userID := c.GetString("userID")

	recommendedEvents, err := h.getRecommendedEventsService.GetRecommendedEvents(userID)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get recommended events", StatusCode: http.StatusInternalServerError})
		return
	}

	trendingEvents, err := h.getTrendingEventsService.GetTrendingEvents()
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get trending events", StatusCode: http.StatusInternalServerError})
		return
	}

	lastParticipatedEvents, err := h.getLastParticipatedEventsService.GetLastParticipatedEvents(userID)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get last participated events", StatusCode: http.StatusInternalServerError})
		return
	}

	subscribedEvents, err := h.getSubscribedEventsService.GetSubscribedEvents(userID)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get subscribed events", StatusCode: http.StatusInternalServerError})
		return
	}

	response := GetEventsResponse{
		RecommendedEvents:     recommendedEvents,
		TrendingEvents:        trendingEvents,
		LastParticipatedEvents: lastParticipatedEvents,
		SubscribedEvents:      subscribedEvents,
	}

	c.JSON(http.StatusOK, response)
}
