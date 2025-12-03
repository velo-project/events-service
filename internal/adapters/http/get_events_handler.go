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
	bucketName                      string
}

func NewGetEventsHandler(
	getRecommendedEventsService ports.GetRecommendedEventsService,
	getTrendingEventsService ports.GetTrendingEventsService,
	getLastParticipatedEventsService ports.GetLastParticipatedEventsService,
	getSubscribedEventsService ports.GetSubscribedEventsService,
	bucketName string,
) *GetEventsHandler {
	return &GetEventsHandler{
		getRecommendedEventsService:     getRecommendedEventsService,
		getTrendingEventsService:        getTrendingEventsService,
		getLastParticipatedEventsService: getLastParticipatedEventsService,
		getSubscribedEventsService:      getSubscribedEventsService,
		bucketName:                      bucketName,
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
	h.setFullImageURLs(recommendedEvents)

	trendingEvents, err := h.getTrendingEventsService.GetTrendingEvents()
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get trending events", StatusCode: http.StatusInternalServerError})
		return
	}
	h.setFullImageURLs(trendingEvents)

	lastParticipatedEvents, err := h.getLastParticipatedEventsService.GetLastParticipatedEvents(userID)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get last participated events", StatusCode: http.StatusInternalServerError})
		return
	}
	h.setFullImageURLs(lastParticipatedEvents)

	subscribedEvents, err := h.getSubscribedEventsService.GetSubscribedEvents(userID)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get subscribed events", StatusCode: http.StatusInternalServerError})
		return
	}
	h.setFullImageURLs(subscribedEvents)

	response := GetEventsResponse{
		RecommendedEvents:     recommendedEvents,
		TrendingEvents:        trendingEvents,
		LastParticipatedEvents: lastParticipatedEvents,
		SubscribedEvents:      subscribedEvents,
	}

	c.JSON(http.StatusOK, response)
}

func (h *GetEventsHandler) setFullImageURLs(events []entities.Event) {
	for i := range events {
		if events[i].ImageURL != nil && *events[i].ImageURL != "" {
			fullURL := h.buildImageURL(*events[i].ImageURL)
			events[i].ImageURL = &fullURL
		}
	}
}

func (h *GetEventsHandler) buildImageURL(objectName string) string {
	return "https://storage.googleapis.com/" + h.bucketName + "/" + objectName
}
