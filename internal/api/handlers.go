package api

import (
	"github.com/ShobenHou/monitor_cp/pkg/kafka"
	"github.com/gin-gonic/gin"
)

func GetAgentConfig(c *gin.Context) {
	agentID := c.Param("id")
	agentConfig, err := getAgentConfigFromDB(agentID) // Replace this with your own logic to get the configuration from the database.
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, agentConfig)
}

func UpdateAgentConfig(c *gin.Context) {
	agentID := c.Param("id")

	var agentConf AgentConf
	err := c.BindJSON(&agentConf)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = updateAgentConfigInDB(agentID, agentConf) // Replace this with your own logic to update the configuration in the database.
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = kafka.PublishConfigToKafka(agentID, agentConf)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Configuration updated successfully."})
}
