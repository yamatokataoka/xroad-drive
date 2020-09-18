package com.yamatokataoka.xroaddrive.api;

import org.springframework.stereotype.Component;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.validation.annotation.Validated;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import javax.validation.constraints.NotBlank;

@ConfigurationProperties(prefix = "xroaddrive.xroad")
@Validated
@Component
// Ignore properties injected by Spring boot
@JsonIgnoreProperties({
  "targetClass",
  "targetSource",
  "advisors",
  "frozen",
  "exposeProxy",
  "preFiltered",
  "proxiedInterfaces",
  "proxyTargetClass"
})
public class XRoadProperties {

  @NotBlank
	private String memberId;
  @NotBlank
  private String commonServiceCode;

  public String getMemberId() {
    return memberId;
  }

  public void setMemberId(String memberId) {
    this.memberId = memberId;
  }

  public String getCommonServiceCode() {
    return commonServiceCode;
  }

  public void setCommonServiceCode(String commonServiceCode) {
    this.commonServiceCode = commonServiceCode;
  }
}
